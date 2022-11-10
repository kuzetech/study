package com.kuze.bigdata.study.junit;

import org.junit.Test;

import java.util.concurrent.TimeUnit;

public class TestTimeout {

    @Test(timeout = 1000) //毫秒
    public void testTimeout() throws InterruptedException {
        TimeUnit.SECONDS.sleep(5); //秒
        System.out.println("in test case 1");
    }

}
