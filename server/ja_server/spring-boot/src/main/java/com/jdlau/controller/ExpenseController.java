package com.jdlau.controller;

import com.jdlau.bean.Expense;
import com.jdlau.service.IExpenseService;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.RequestMapping;

@RestController
public class ExpenseController {

    @Autowired
    IExpenseService expenseService;

    @RequestMapping("/showExpenses")
    public String findExpenses(Model model) {

        List<Expense> expenses = (List<Expense>) expenseService.findAll();

        System.out.println(expenses);
        model.addAttribute("expenses", expenses);

        return "showExpenses";
    }
}