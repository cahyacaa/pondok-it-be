package main

import "fmt"

type Invoice struct {
    ID    int
    Items []string
}

func (inv *Invoice) Total() float64 {
    // hitung total harga
    return 100.0
}

// printer.go
package billing

import "fmt"

func PrintInvoice(inv *Invoice) {
    fmt.Printf("Invoice #%d\nItems: %v\nTotal: %.2f\n",
        inv.ID, inv.Items, inv.Total())
}
