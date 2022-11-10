package com.kuze.bigdata.study.junit;

import org.junit.runner.RunWith;
import org.junit.runners.Suite;

@RunWith(Suite.class)
@Suite.SuiteClasses({
        /**
         * 这里的配置影响程序运行的顺序
         */
        TestParameterized.class,
        TestException.class
})
public class TestSuite {}
