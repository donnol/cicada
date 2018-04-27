package com.jdlau.service;

import com.jdlau.bean.Expense;
import com.jdlau.repository.ExpenseRepository;
import java.util.List;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

@Service
public class ExpenseService implements IExpenseService {

    @Autowired
    private ExpenseRepository repository;

    @Override
    public List<Expense> findAll() {

        List<Expense> expenses = (List<Expense>) repository.findAll();

        return expenses;
    }
}