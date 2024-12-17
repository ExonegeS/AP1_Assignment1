package ports

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/ExonegeS/AP1_Assignment1/internal/shapes/domain"
	"github.com/ExonegeS/AP1_Assignment1/internal/shapes/service"
)

func NewEndpointPrintShapes(s service.Service, logger *slog.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte("Check the Stdout!"))

		logger.Error("PrintShapes recieved successfully")

		shapes := make([]domain.Shape, 0)
		shapes = append(shapes, &domain.Circle{
			Radius: 15.908,
		})
		shapes = append(shapes, &domain.Square{
			Length: 15,
		})
		shapes = append(shapes, &domain.Rectangle{
			Width:  31.0009,
			Length: 13.63,
		})
		shapes = append(shapes, &domain.Square{
			Length: 13.7,
		})
		shapes = append(shapes, &domain.Square{
			Length: 5,
		})
		shapes = append(shapes, &domain.Triangle{
			SideA: 6,
			SideB: 8,
			SideC: 10,
		})
		shapes = append(shapes, &domain.Triangle{
			SideA: 6,
			SideB: 8,
			SideC: 5,
		})
		shapes = append(shapes, &domain.Circle{
			Radius: 15.908,
		})

		for _, shape := range shapes {
			PrintShapeDetails(shape)
		}
	}
}

func PrintShapeDetails(s domain.Shape) {
	fmt.Printf("Shape: %#v\n", s)
	fmt.Printf("Area: %v | Perimeter: %v \n", s.Area(), s.Perimeter())
}
