package accounts

import (
	"fmt"
	"hse_mini_course/accounts/dto"
	"hse_mini_course/accounts/models"
	"math/rand/v2"
	"sync"

	"github.com/gofiber/fiber/v2"
)

const (
	IDS_FROM    = 100000
	IDS_TO      = 999999
	GUARD_LIMIT = 20
)

func generateId() uint32 {
	return rand.Uint32N(IDS_TO - IDS_FROM + 1) + IDS_FROM
}

type Handler struct {
	accounts      map[string]*models.Account
	ids           map[uint32]bool
	accountsGuard *sync.RWMutex
	idsGuard      *sync.RWMutex
}

func NewHandler() *Handler {
	return &Handler{
		accounts:      make(map[string]*models.Account),
		ids:           make(map[uint32]bool),
		accountsGuard: &sync.RWMutex{},
		idsGuard:      &sync.RWMutex{},
	}
}

func (h *Handler) validateName(ctx *fiber.Ctx, name string) (error, bool) {
	if len(name) == 0 {
		return ctx.Status(fiber.StatusForbidden).SendString("name cannot be empty"), false
	}
	if name == "new" {
		return ctx.Status(fiber.StatusForbidden).SendString("name cannot be \"new\""), false
	}
	for _, ch := range name {
		if !(('a' <= ch && ch <= 'z') || ('A' <= ch && ch <= 'Z')) {
			return ctx.Status(fiber.StatusForbidden).SendString(
				fmt.Sprintf("name \"%s\" must consist of only latin letters", name),
			), false
		}
	}
	return nil, true
}

func (h *Handler) CreateAccount(ctx *fiber.Ctx) error {
	var request dto.CreateAccountRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid request")
	}
	name := request.Name
	err, valid := h.validateName(ctx, name)
	if !valid {
		return err
	}

	h.accountsGuard.Lock()

	_, taken := h.accounts[name]
	if taken {
		h.accountsGuard.Unlock()
		return ctx.Status(fiber.StatusForbidden).SendString(
			fmt.Sprintf("name \"%s\" already taken", name),
		)
	}

	new_acc := models.Account{
		Id:      generateId(),
		Name:    name,
		Balance: 0,
	}

	var iter_limit int
	h.idsGuard.Lock()

	for iter_limit = 0; iter_limit < GUARD_LIMIT && h.ids[new_acc.Id]; iter_limit++ {
		new_acc.Id = generateId()
	}
	if iter_limit == GUARD_LIMIT {
		h.idsGuard.Unlock()
		h.accountsGuard.Unlock()
		ctx.Context().Logger().Printf("[ERROR]: Hit iteration limit while generating ids\n")
		return ctx.Status(fiber.StatusInternalServerError).Send(nil)
	}
	h.ids[new_acc.Id] = true
	h.accounts[name] = &new_acc

	h.idsGuard.Unlock()
	h.accountsGuard.Unlock()
	return ctx.Status(fiber.StatusCreated).JSON(dto.AccountToResponse(new_acc))
}

func (h *Handler) GetUser(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	h.accountsGuard.RLock()

	account, found := h.accounts[name]
	if !found {
		h.accountsGuard.RUnlock()
		return ctx.Status(fiber.StatusNotFound).SendString(
			fmt.Sprintf("name \"%s\" not found", name),
		)
	}
	h.accountsGuard.RUnlock()
	return ctx.Status(fiber.StatusOK).JSON(dto.AccountToResponse(*account))
}

func (h *Handler) NewTransaction(ctx *fiber.Ctx) error {
	name := ctx.Params("name")
	var request dto.NewTransactionRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid request")
	}
	delta := request.Delta

	h.accountsGuard.Lock()

	acc, found := h.accounts[name]
	if !found {
		h.accountsGuard.Unlock()
		return ctx.Status(fiber.StatusNotFound).SendString(
			fmt.Sprintf("name \"%s\" not found", name),
		)
	}
	acc.Balance += delta
	h.accountsGuard.Unlock()
	return ctx.Status(fiber.StatusOK).JSON(dto.AccountToResponse(*acc))
}

func (h *Handler) ChangeName(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	var request dto.ChangeNameRequest
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).SendString("invalid request")
	}
	newName := request.NewName
	err, valid := h.validateName(ctx, newName)
	if !valid {
		return err
	}

	h.accountsGuard.Lock()
	acc, found := h.accounts[name]
	if !found {
		h.accountsGuard.Unlock()
		return ctx.Status(fiber.StatusNotFound).SendString(
			fmt.Sprintf("name \"%s\" not found", name),
		)
	}
	_, found = h.accounts[newName]
	if found {
		h.accountsGuard.Unlock()
		return ctx.Status(fiber.StatusForbidden).SendString(
			fmt.Sprintf("name \"%s\" already taken", newName),
		)
	}

	delete(h.accounts, name)
	acc.Name = newName
	h.accounts[newName] = acc
	h.accountsGuard.Unlock()

	return ctx.Status(fiber.StatusOK).JSON(dto.AccountToResponse(*acc))
}

func (h *Handler) DeleteAccount(ctx *fiber.Ctx) error {
	name := ctx.Params("name")

	h.accountsGuard.Lock()
	account, found := h.accounts[name]
	if !found {
		h.accountsGuard.Unlock()
		return ctx.Status(fiber.StatusNotFound).SendString(
			fmt.Sprintf("name \"%s\" not found", name),
		)
	}
	id := account.Id

	h.idsGuard.Lock()
	delete(h.ids, id)
	h.idsGuard.Unlock()

	delete(h.accounts, name)
	h.accountsGuard.Unlock()

	return ctx.Status(fiber.StatusOK).Send(nil)
}
