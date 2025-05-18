package main

import (
	"fmt"
	"school_schedule_2/internal/adapter/in_memory_storage/office_storage"
	"school_schedule_2/internal/usecase/office_usecase"
)

func main() {
	officeRepo := office_storage.NewOfficeRepo()

	officeUseCase := office_usecase.NewOfficeUseCase(officeRepo)

	createdOffice, err := officeUseCase.CreateOffice(office_usecase.CreateOfficeReq{
		Name: 10,
	})
	if err != nil {
		fmt.Println(err)
	}
	createdOffice2, err := officeUseCase.CreateOffice(office_usecase.CreateOfficeReq{
		Name: 21,
	})
	if err != nil {
		fmt.Println(err)
	}
	createdOffice3, err := officeUseCase.CreateOffice(office_usecase.CreateOfficeReq{
		Name: 22,
	})
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(createdOffice, createdOffice2, createdOffice3)

}
