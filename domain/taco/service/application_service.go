package service

import (
	"github.com/fwchen/jellyfish/domain/taco"
	"github.com/fwchen/jellyfish/domain/taco/command"
	"github.com/fwchen/jellyfish/domain/taco/factory"
	"github.com/fwchen/jellyfish/domain/taco/repository"
	"github.com/fwchen/jellyfish/domain/taco_box"
	"github.com/fwchen/jellyfish/domain/taco_box/service"
	"github.com/fwchen/jellyfish/util"
	"github.com/juju/errors"
)

func NewTacoApplicationService(tacoRepo repository.Repository, tacoBoxPermissionService *service.TacoBoxPermissionService) *TacoApplicationService {
	return &TacoApplicationService{
		tacoRepo:                 tacoRepo,
		tacoBoxPermissionService: tacoBoxPermissionService,
	}
}

type TacoApplicationService struct {
	tacoRepo                 repository.Repository
	tacoBoxPermissionService *service.TacoBoxPermissionService
}

func (t *TacoApplicationService) GetTacos(userID string, status []taco.Status, boxName string) ([]taco.Taco, error) {
	var tacoTypeStr string
	var boxId *string = nil
	// TODO 放在前端做
	if taco_box.ContainTypeTacoBox(boxName) {
		tacoTypeStr = boxName
	} else if boxName == taco_box.TacoBoxAll {
		tacoTypeStr = string(taco.Task)
	} else {
		tacoTypeStr = string(taco.Task)
		if boxName == "" {
			boxId = nil
		} else {
			boxId = util.PointerStr(boxName)
		}
	}
	tacoType := taco.Type(tacoTypeStr)
	return t.tacoRepo.List(userID, taco.ListTacoFilter{
		Statues: status,
		Type:    &tacoType,
		BoxId:   boxId,
	})
}

func (t *TacoApplicationService) CreateTaco(command *command.CreateTacoCommand, userId string) (*string, error) {
	isInBox := command.BoxId != nil
	var maxOrder *float64
	var err error
	if isInBox {
		maxOrder, err = t.tacoRepo.MaxOrderByBoxId(userId)
	} else {
		maxOrder, err = t.tacoRepo.MaxOrderByCreatorId(userId)

	}
	if err != nil {
		return nil, errors.Trace(err)
	}
	command.Order = *maxOrder + float64(10)
	if command.BoxId != nil {
		if !taco_box.ContainCommonTacoBox(*command.BoxId) {
			can, err := t.tacoBoxPermissionService.CheckUserCanOperation(*command.BoxId, userId)
			if err != nil {
				return nil, errors.Trace(err)
			}
			if !can {
				return nil, errors.Forbiddenf("user [userId = %s] forbidden create taco in box [boxId = %s]", userId, *command.BoxId)
			}
		} else {
			command.BoxId = nil
		}
	}
	ta := factory.NewTacoFromCreateCommand(command, userId)
	return t.tacoRepo.Save(ta)
}

func (t *TacoApplicationService) UpdateTaco(command command.UpdateTacoCommand) error {
	ta, err := t.tacoRepo.FindById(command.TacoId)
	if err != nil {
		return errors.Trace(err)
	}
	ta.Content = command.Content
	ta.Detail = command.Detail
	ta.Deadline = command.Deadline
	ta.Status = command.Status
	ta.BoxId = command.BoxId
	_, err = t.tacoRepo.Save(ta)
	return err
}

func (t *TacoApplicationService) DeleteTaco(id string) error {
	return t.tacoRepo.Delete(id)
}

func (t *TacoApplicationService) Reorder(command *command.SortTacoCommand, userId string) error {
	tacoType := taco.Type("Task")
	status := []taco.Status{taco.Status("Doing")}
	tacos, err := t.tacoRepo.List(userId, taco.ListTacoFilter{
		Statues: status,
		Type:    &tacoType,
		BoxId:   nil, // TODO
	})
	tacos = taco.SortTacosByOrder(tacos)
	if err != nil {
		return errors.Trace(err)
	}
	if len(tacos) == 0 {
		return errors.Errorf("todos is empty")
	}
	moveTacoIndex := taco.IndexOfTacos(tacos, command.TacoId)
	targetTacoIndex := taco.IndexOfTacos(tacos, command.TargetTacoId)
	tacos = SortTacos(tacos, moveTacoIndex, targetTacoIndex)
	for i := 0; i < len(tacos); i++ {
		tacos[i].Order = float64(i * 10)
	}
	err = t.tacoRepo.SaveList(tacos)
	if err != nil {
		return errors.Trace(err)
	}
	return nil
}

func SortTacos(tacos []taco.Taco, moveTacoIndex int, targetTacoIndex int) []taco.Taco {
	moveTaco := tacos[moveTacoIndex]
	nTacos := taco.SliceRemove(tacos, moveTacoIndex)
	nTacos = taco.InsertInTacos(nTacos, moveTaco, targetTacoIndex)
	return nTacos
}
