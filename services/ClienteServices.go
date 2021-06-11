package services

import (
	"github.com/OakAnderson/academia/data/model"
	"github.com/OakAnderson/academia/repository"
)

func findAll() []model.Cliente {
	clientes := repository.FindAll()
	if clientes == nil {
		panic("Não há nenhum registro no banco de dados.")
	}
	return clientes
}

func MediaPeso() float32 {
	var media float32
	clientes := findAll()
	for _, cliente := range clientes {
		media += cliente.PesoAtual
	}
	return media/float32(len(clientes))
}

func ClienteQueMaisPerdeuPeso() model.Cliente{
	var peso float32
	var cliente model.Cliente
	for _, c := range findAll() {
		if c.PesoInicial - c.PesoAtual > peso {
			cliente = c
		}
	}

	return cliente
}
