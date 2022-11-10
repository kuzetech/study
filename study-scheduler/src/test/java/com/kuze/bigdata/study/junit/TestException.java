package com.kuze.bigdata.study.junit;

import org.junit.Test;

public class TestException {

    @Test(expected = ArithmeticException.class)
    public void testException() {
        System.out.println("执行 testException 方法");
        int a = 0;
        int b = 1 / a;
    }

}
