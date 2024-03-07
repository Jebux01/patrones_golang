package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	paymentProcessor := NewPaymentProcessor()

	cashPayment := &CashPayment{}
	paymentProcessor.RegisterPaymentMethod("cash", cashPayment)

	cardPayment := &CardPayment{}
	paymentProcessor.RegisterPaymentMethod("card", cardPayment)

	psePayment := &PSEPayment{}
	paymentProcessor.RegisterPaymentMethod("pse", psePayment)

	paymentLogger := &PaymentLogger{}
	paymentProcessor.Attach(paymentLogger)

	for {
		messageInitial()
		clearScreen()

		fmt.Println("Seleccione una opción:")
		fmt.Println("1. Pagar")
		fmt.Println("2. Deshacer el pago")
		fmt.Println("3. Ver registros")
		fmt.Println("4. Salir")

		option, err := getOption(reader)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		if option == 4 {
			fmt.Println("Saliendo...")
			break
		}

		if option == 3 {
			showRecords()
			reader.ReadString('\n')
			continue
		}

		amount, err := getAmount(reader)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}

		paymentOption, err_bool, method := selectPaymentOption(reader, paymentProcessor)
		if !err_bool {
			fmt.Println("Error:", err)
			continue
		}

		switch option {
		case 1:
			clearScreen()
			discount, err_d := selectDiscount(reader)
			if err_d != nil {
				fmt.Println("Error:", err)
				continue
			}
			methodWithDiscount := NewPaymentDecorator(paymentOption, float32(discount))
			fmt.Printf("Ha seleccionado pagar con %s\n", method)
			fmt.Printf("Descuento aplicado: %v%%\n", discount)
			fmt.Println("Por favor espere...")
			paymentProcessor.Execute(methodWithDiscount, map[string]interface{}{"amount": amount, "name": method})
			fmt.Println("\n Presione enter para continuar...")
			reader.ReadString('\n')
			continue
		case 2:
			clearScreen()
			uuid, err := getUUID(reader)
			if err != nil {
				fmt.Println("Error:", err)
				continue
			}
			paymentProcessor.Detach(paymentOption, map[string]interface{}{"amount": amount, "name": method, "uuid": uuid})
			time.Sleep(2 * time.Second)
			clearScreen()
			continue
		default:
			fmt.Println("Opción no válida")
		}
	}
}

func showRecords() {
	clearScreen()
	fmt.Println("Registros:")
	for _, reg := range getAllRegs() {
		fmt.Println(reg)
	}
	fmt.Println("Presione enter para continuar...")
}

func messageInitial() {
	clearScreen()
	text := `
Sacramentum nocturnarum nefarius
Sacramentum nocturnarum nefarius
Oremus deus sanctus, deum filium
Dominum martyrum
Oremus convertere apostolicus
Cedere animus debitus
Et Catholicus debere
Deum animalum dominum`

	for _, char := range text {
		fmt.Print(string(char))
		time.Sleep(10 * time.Millisecond)
	}

	fmt.Println()
}

func clearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func getOption(reader *bufio.Reader) (int, error) {
	fmt.Print("Ingrese su opción: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	option, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		return 0, err
	}
	return option, nil
}

func getUUID(reader *bufio.Reader) (string, error) {
	fmt.Print("Ingrese el UUID: ")
	uuid, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	return uuid[:len(uuid)-1], nil
}

func getAmount(reader *bufio.Reader) (float32, error) {
	fmt.Print("Ingrese el monto: ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	amount, err := strconv.ParseFloat(input[:len(input)-1], 32)
	if err != nil {
		return 0, err
	}
	return float32(amount), nil
}

func selectPaymentOption(reader *bufio.Reader, paymentProcesor *PaymentProcessor) (Payment, bool, string) {
	fmt.Println("Seleccione el método de pago:")
	fmt.Println("1. Tarjeta")
	fmt.Println("2. Efectivo")
	fmt.Println("3. PSE")

	option, err := getOption(reader)
	if err != nil {
		panic(err)
	}

	switch option {
	case 1:
		payment, repo := paymentProcesor.GetPaymentMethod("card")
		return payment, repo, "Tarjeta"
	case 2:
		payment, repo := paymentProcesor.GetPaymentMethod("cash")
		return payment, repo, "Efectivo"
	case 3:
		payment, repo := paymentProcesor.GetPaymentMethod("pse")
		return payment, repo, "PSE"
	default:
		panic("Opción no válida")
	}
}

func selectDiscount(reader *bufio.Reader) (int, error) {
	fmt.Print("¿Desea agregar un descuento? (s/n): ")
	input, err := reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	if input[:len(input)-1] != "s" {
		return 0, nil
	}

	fmt.Print("Ingrese el descuento en porcentaje: ")
	input, err = reader.ReadString('\n')
	if err != nil {
		return 0, err
	}
	discount, err := strconv.Atoi(input[:len(input)-1])
	if err != nil {
		return 0, err
	}
	return discount, nil
}
