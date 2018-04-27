package com.jdlau.service;

import com.jdlau.bean.Expense;
import java.util.List;

public interface IExpenseService {

    public List<Expense> findAll();
}