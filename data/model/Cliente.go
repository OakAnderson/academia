package model

import (
	"fmt"
	"github.com/OakAnderson/academia/data/model/interfaces"
)

type Cliente struct {
	Id int
	Cpf string
	Sexo string
	PesoInicial float32
	PesoAtual float32
	Altura float32
	Biceps float32
	Peito float32
}

func (c Cliente) String() string {
	return fmt.Sprintf("  ID - %d\n  CPF - %s\n  Sexo - %s\n  Peso Inicial - %.2f\n  Peso Atual - %.2f\n  Altura - %.2f\n  Bíceps - %.2f\n  Peito - %.2f\n",
		c.Id, c.formatCPF(), c.formatSexo(), c.PesoInicial, c.PesoAtual, c.Altura, c.Biceps, c.Peito)
}

func (c Cliente) formatCPF() string {
	return c.Cpf[0:3]+"."+c.Cpf[3:6]+"."+c.Cpf[6:9]+"-"+c.Cpf[9:11]
}

func (c Cliente) formatSexo() string {
	if c.Sexo == "m"{
		return "Masculino"
	}
	return "Feminino"
}

func (c *Cliente) Scan(row interfaces.SqlScan) error {
	return row.Scan(
		&c.Id, &c.Cpf, &c.Sexo,
		&c.PesoInicial, &c.PesoAtual, &c.Altura,
		&c.Biceps, &c.Peito)
}
