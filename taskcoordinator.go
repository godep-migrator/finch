package main

import "time"

type TaskCoordinator struct {
	storage MetaTaskStore
}

func (tc *TaskCoordinator) Add(desc string) (*Task, error) {
	t := &Task{
		Desc:   desc,
		Active: time.Now(),
	}

	err := tc.storage.SaveTask(t)
	return t, err
}

func (tc *TaskCoordinator) Delay(id string, until time.Time) error {
	t, err := tc.storage.GetTask(id)
	if err != nil {
		return err
	}

	t.Active = until
	err = tc.storage.SaveTask(t)
	if err != nil {
		return err
	}

	return nil
}

func (tc *TaskCoordinator) Select(ids ...string) error {
	// make sure we have all the tasks before we perform the operation
	tasks := []*Task{}
	for _, id := range ids {
		t, err := tc.storage.GetTask(id)
		if err != nil {
			return err
		}

		tasks = append(tasks, t)
	}

	// select all the tasks and save them to the DB
	for _, t := range tasks {
		t.Selected = true
	}
	tc.storage.SaveTask(tasks...)

	return nil
}
