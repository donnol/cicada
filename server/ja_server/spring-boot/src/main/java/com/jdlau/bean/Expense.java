package com.jdlau.bean;

import java.sql.Time;
import java.util.Objects;
import javax.persistence.Entity;
import javax.persistence.GeneratedValue;
import javax.persistence.GenerationType;
import javax.persistence.Id;
import javax.persistence.Table;

@Entity
@Table(name = "T_EXPENSE")
public class Expense {

    @Id
    @GeneratedValue(strategy = GenerationType.AUTO)
    private Long id;

    private Long user_id;
    private String user_name;
    private float pay;
    private String thing;
    private Time created_at;
    private String created_on;
    private Time updated_at;

    public Expense() {
    }

    public Long getId() {
        return id;
    }

    public void setId(Long id) {
        this.id = id;
    }

    public Long getUserID() {
        return user_id;
    }

    public void setUserID(Long userID) {
        this.user_id = userID;
    }

    public String getUserName() {
        return user_name;
    }

    public void setUserName(String name) {
        this.user_name = name;
    }

    public float getPay() {
        return pay;
    }

    public void setPay(float pay) {
        this.pay = pay;
    }

    public String getThing() {
        return thing;
    }

    public void setThing(String thing) {
        this.thing = thing;
    }

    public String getCreatedOn() {
        return created_on;
    }

    public void setCreatedOn(String co) {
        this.created_on = co;
    }

    public Time getCreatedAt() {
        return created_at;
    }

    public void setCreatedAt(Time ca) {
        this.created_at = ca;
    }

    public Time getUpdatedAt() {
        return updated_at;
    }

    public void setUpdatedAt(Time ua) {
        this.updated_at = ua;
    }
}