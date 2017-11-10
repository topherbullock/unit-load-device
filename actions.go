package main

import "code.cloudfoundry.org/garden"

type UnitLoadDevice struct {
	GardenClient garden.Client
}

type Task struct{}

func (t *Task) containerSpec() garden.ContainerSpec {
	return garden.ContainerSpec{}
}

func (t *Task) processSpec() garden.ProcessSpec {
	return garden.ProcessSpec{}
}

func (uld *UnitLoadDevice) RunTask(task Task) error {
	container, err := uld.GardenClient.Create(task.containerSpec())
	if err != nil {
		return err
	}

	container.Run(task.processSpec(), garden.ProcessIO{})

	return nil
}
