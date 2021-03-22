package Contas

import c "Banco/clientes"

type ContaCorrente struct {
	Cliente    c.Titular
	NumAgencia int
	NumConta   int
	saldo      float64
}

func (c *ContaCorrente) Sacar(valorSaque float64) string {
	podeSacar := valorSaque > 0 && valorSaque <= c.saldo
	if podeSacar {
		c.saldo -= valorSaque
		return "Saque realizado com sucesso!"
	} else {
		return "Saque não efetuado!"
	}
}

func (c *ContaCorrente) Depositar(valorDeposito float64) (string, float64) {
	if valorDeposito > 0 {
		c.saldo += valorDeposito
		return "Deposito realizado com sucesso!", c.saldo
	} else {
		return "Deposito não efetuado!", c.saldo
	}
}

func (c *ContaCorrente) Transferir(valorTransferencia float64, Destino *ContaCorrente) bool {
	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		Destino.Depositar(valorTransferencia)
		c.saldo -= valorTransferencia
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {
	return c.saldo
}
