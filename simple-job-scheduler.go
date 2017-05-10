package main

//
// Defines a simple job scheduler that runs one (ideally short) job at a time.
// Each job has a current or future RunAt time when it will be run.
//
// Users of this implement Job, JobGetter, and JobRunner
//

import (
	"fmt"
	"github.com/beeker1121/goque"
	"github.com/pkg/errors"
	"log"
	"sort"
	"sync"
	"time"
)

const (
	JobSuccess      = JobStatus("success")
	JobReady        = JobStatus("ready")
	JobError        = JobStatus("error")
	intervalSeconds = 5
)

type JobStatus string

// ==========================================================
// Job, JobGetter and JobRunner are all implemented by the user of this library
// ==========================================================

// Job interface is implemented by the user to represent the job he wants us to run.
type Job interface {
	GetId() uint64
	SetId(id uint64)
	GetStatus() JobStatus
	SetStatus(status JobStatus)
	GetRunAt() time.Time
}

// JobGetter is implemented by the user to retrieve the Job instance from a goque.Item.
// This is less than idea becuase the user shouldn't know we are using goque.
type JobGetter func(item *goque.Item) (Job, error)

// JobRunner is implemented by the user to run their job
type JobRunner func(job Job) (JobStatus, error)

//
// Jobs is a sortable slice of Job instances
//
type Jobs []Job

func (slice Jobs) Len() int {
	return len(slice)
}

func (slice Jobs) Less(i, j int) bool {
	return slice[i].GetRunAt().Before(slice[j].GetRunAt())
}

func (slice Jobs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

//
// JobScheduler manages the jobs that need to be run
//
type JobScheduler struct {
	Name      string
	getter    JobGetter
	runner    JobRunner
	readyJobs Jobs
	mutex     sync.Mutex // control access to readyJobs
}

func NewJobScheduler(name string, getter JobGetter, runner JobRunner) *JobScheduler {
	job := &JobScheduler{Name: name, getter: getter, runner: runner}
	err := job.initialize()
	if err != nil {
		log.Printf("Error initializing JobScheduler %s: %v", name, err)
		return nil
	}
	return job
}

// initialize adds any persistent jobs so they can be run
func (js *JobScheduler) initialize() error {

	js.mutex.Lock()
	defer js.mutex.Unlock()

	// Collect and sort all persisted "ready" jobs
	q, err := goque.OpenQueue(js.Name)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error opening queue %s", js.Name))
	}
	defer q.Close()

	// Get each persisted job and add it to the readyJobs field
	for i := uint64(0); i < q.Length(); i++ {
		item, err := q.PeekByOffset(i)
		if err != nil {
			return errors.Wrap(err, "Error in inititialize calling PeekByOffset")
		}
		job, err := js.getter(item)
		if err != nil {
			return errors.Wrap(err, "Error in inititialize calling jobGetter")
		}
		if job.GetStatus() == JobReady {
			job.SetId(item.ID)
			js.readyJobs = append(js.readyJobs, job)
		} else {
			log.Println("Initialize is ignoring job:", job)
		}

	}

	// This goroutine regularly runs the available jobs one by one
	go func() {
		log.Println("Started goroutine")
		for {
			time.Sleep(time.Duration(intervalSeconds) * time.Second)
			log.Println("Woke up...")
			if len(js.readyJobs) > 0 {
				js.checkJobs()
			}
		}
	}()

	return nil
}

// AddJob persists a new job to be run
func (js *JobScheduler) AddJob(job Job) error {

	js.mutex.Lock()
	defer js.mutex.Unlock() // Persist the job
	q, err := goque.OpenQueue(js.Name)
	if err != nil {
		return errors.Wrap(err, fmt.Sprintf("Error opening queue %s", js.Name))
	}
	defer q.Close()

	// Set the initial state before we persist it
	job.SetStatus(JobReady)

	item, err := q.EnqueueObject(job)
	if err != nil {
		return errors.Wrap(err, "Error adding job to work queue")
	}
	job.SetId(item.ID)

	// Add job to our in-memory list and select the next job to run
	js.readyJobs = append(js.readyJobs, job)

	log.Println("job added: ", item.ID, job)

	return nil
}

// checkJobs is called every minute (or so) to see if there is any work to do
func (js *JobScheduler) checkJobs() {

	// Lock for the duration of this method
	js.mutex.Lock()
	defer js.mutex.Unlock()

	log.Printf("Checking %d job(s)\n", len(js.readyJobs))
	if len(js.readyJobs) == 0 {
		return
	}

	// Sort the jobs by ascending RunAt time
	sort.Sort(js.readyJobs)

	// Loop through each job in the list
	for _, job := range js.readyJobs {
		// Run the job only if it is time to run
		if job.GetRunAt().Before(time.Now()) {

			// Run the job
			jobStatus, err := js.runner(job)
			if err != nil {
				log.Printf("Error running jobId %d in queue %s    %v\n", job.GetId(), js.Name, err)
				jobStatus = JobError
			}

			// Remove the first one (the one we just ran)
			js.readyJobs = js.readyJobs[1:]

			// Update persistent version of job with new status
			job.SetStatus(jobStatus)
			err = updateJob(js.Name, job)
			if err != nil {
				log.Println("Error updating jobId %d in queue %s    %v\n", job.GetId(), js.Name, err)
				return
			}

		} else {
			log.Printf("JobId %d is not ready to run until %v\n", job.GetId(), job.GetRunAt())
		}
	}
}

// updateJob puts the new job state back into the database
func updateJob(queueName string, job Job) error {
	q, err := goque.OpenQueue(queueName)
	if err != nil {
		return errors.Wrap(err, "Error opening queue:")
	}
	defer q.Close()

	_, err = q.UpdateObject(job.GetId(), job)
	if err != nil {
		return errors.Wrap(err, "Error updating job in queue:")
	}

	log.Println("Updated job with new status:", job)
	return nil
}

// ===========================================================================
//
// Everything above is the "library" and not to be modified by the user
// Everything below is the user implementation
//
// ===========================================================================

func main() {
	js := NewJobScheduler("sleep", SleepJobGetter, SleepJobRunner)
	err := js.AddJob(&SleepJob{SleepSeconds: 3, Status: JobReady, RunAt: time.Now().Add(time.Duration(20) * time.Second)})
	checkErr(err)
	time.Sleep(time.Duration(5) * time.Second)
	err = js.AddJob(&SleepJob{SleepSeconds: 7, Status: JobReady, RunAt: time.Now()})
	checkErr(err)
	time.Sleep(time.Duration(2) * time.Second)
	err = js.AddJob(&SleepJob{SleepSeconds: 1})
	checkErr(err)
	time.Sleep(time.Duration(30) * time.Second)
	log.Println("Done")
}

func checkErr(err error) {
	if err != nil {
		log.Fatalln("Error adding Job:", err)
	}
}

// SleepJob requires public fields and public getters/setters.
// The methods on SleepJob implement the Job interface
type SleepJob struct {
	Id           uint64
	Status       JobStatus
	RunAt        time.Time
	SleepSeconds int64
}

func (my *SleepJob) GetId() uint64 {
	return my.Id
}

func (my *SleepJob) SetId(id uint64) {
	my.Id = id
}

func (my *SleepJob) GetStatus() JobStatus {
	return my.Status
}

func (my *SleepJob) SetStatus(status JobStatus) {
	my.Status = status
}

func (my *SleepJob) GetRunAt() time.Time {
	return my.RunAt
}

// SleepJobGetter decodes the SleepJob from the goque Item value.
// This is less than idea becuase the user shouldn't know we are using goque.
// It's really just boilerplate code to specify the type of job stored in the queue
// The gob package requires the exact type when decoding
func SleepJobGetter(item *goque.Item) (Job, error) {
	var job SleepJob
	err := item.ToObject(&job)
	if err != nil {
		return &job, errors.Wrap(err, "Error getting job from item")
	}
	return &job, nil
}

// SleepJobRunner runs the sleep job
func SleepJobRunner(job Job) (JobStatus, error) {

	// Ensure it's really a SleepJob
	sleepJob, ok := job.(*SleepJob)
	if !ok {
		return JobError, errors.New("Not running a SleepJob")
	}

	log.Printf("SleepJob %d is sleeping for %d seconds", sleepJob.Id, sleepJob.SleepSeconds)
	time.Sleep(time.Duration(sleepJob.SleepSeconds) * time.Second)

	log.Printf("Job %d sleeping for %d seconds", sleepJob.Id, sleepJob.SleepSeconds)
	return JobSuccess, nil
}
