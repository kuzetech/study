package com.kuze.bigdata.study.dolphin.api;

import com.alibaba.fastjson.JSONArray;
import com.alibaba.fastjson.JSONObject;
import org.apache.hc.client5.http.classic.methods.HttpGet;
import org.apache.hc.client5.http.classic.methods.HttpPost;
import org.apache.hc.client5.http.entity.UrlEncodedFormEntity;
import org.apache.hc.client5.http.impl.classic.CloseableHttpClient;
import org.apache.hc.client5.http.impl.classic.CloseableHttpResponse;
import org.apache.hc.client5.http.impl.classic.HttpClients;
import org.apache.hc.core5.http.NameValuePair;
import org.apache.hc.core5.http.ParseException;
import org.apache.hc.core5.http.io.entity.EntityUtils;
import org.apache.hc.core5.http.message.BasicNameValuePair;
import org.junit.After;
import org.junit.Before;
import org.junit.Test;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

public class TestProjects {

    private final static Logger logger = LoggerFactory.getLogger(TestProjects.class);

    public final static String API_CONNECT_URL = "http://localhost:12345/dolphinscheduler";
    public final static String USER_TOKEN_KEY = "token";
    public final static String USER_TOKEN = "2a3c456d82c17f580ff4dccc58f8c295";

    private CloseableHttpClient httpclient;

    @Before
    public void before() {
        httpclient = HttpClients.createDefault();
    }

    @After
    public void after() throws IOException {
        httpclient.close();
    }


    @Test
    public void testGetProjectsList() throws IOException, ParseException {
        HttpGet request = new HttpGet(API_CONNECT_URL + "/projects/created-and-authed");
        request.setHeader(USER_TOKEN_KEY, USER_TOKEN);
        CloseableHttpResponse response = httpclient.execute(request);
        String content = EntityUtils.toString(response.getEntity(), "UTF-8");
        JSONObject jsonObject = JSONObject.parseObject(content);
        JSONArray dataArray = jsonObject.getJSONArray("data");
        for (Object o : dataArray) {
            JSONObject x = (JSONObject) o;
            String name = x.getString("name");
            Long code = x.getLong("code");
            System.out.println("项目CODE为：" + code + "，项目名为：" + name);
        }
        response.close();
    }

    @Test
    public void testCreateProject() throws IOException, ParseException {
        HttpPost request = new HttpPost(API_CONNECT_URL + "/projects");
        request.setHeader(USER_TOKEN_KEY, USER_TOKEN);

        // set parameters
        List<NameValuePair> parameters = new ArrayList<>();
        parameters.add(new BasicNameValuePair("projectName", "qwe"));
        parameters.add(new BasicNameValuePair("description", "qwe"));
        UrlEncodedFormEntity formEntity = new UrlEncodedFormEntity(parameters);
        request.setEntity(formEntity);

        CloseableHttpResponse response = httpclient.execute(request);
        String content = EntityUtils.toString(response.getEntity(), "UTF-8");
        JSONObject jsonObject = JSONObject.parseObject(content);
        JSONObject data = jsonObject.getJSONObject("data");
        Long code = data.getLong("code");
        String createProjectName = data.getString("name");

        System.out.println("创建的项目名为：" + createProjectName);
        System.out.println("创建的项目CODE为：" + code);
        response.close();
    }




}
