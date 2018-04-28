package com.jdlau.controller;

import com.jdlau.bean.Expense;
import com.jdlau.service.IExpenseService;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.ResponseBody;

@RestController
public class ExpenseController {

    @Autowired
    IExpenseService expenseService;

    @RequestMapping(value = "/showExpenses", produces = "application/json")
    public @ResponseBody List<Expense> findExpenses() {

        List<Expense> expenses = (List<Expense>) expenseService.findAll();

        return expenses;
    }
}