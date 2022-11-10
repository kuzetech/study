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

import java.io.IOException;
import java.util.ArrayList;
import java.util.List;

import static com.kuze.bigdata.study.dolphin.api.TestProjects.*;

public class TestProcessDefinition {

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
    public void testGetProcessDefinitionList() throws IOException, ParseException {
        HttpGet request = new HttpGet(API_CONNECT_URL + "/projects/6564650887104/process-definition/query-process-definition-list");
        request.setHeader(USER_TOKEN_KEY, USER_TOKEN);
        CloseableHttpResponse response = httpclient.execute(request);
        String content = EntityUtils.toString(response.getEntity(), "UTF-8");
        JSONObject jsonObject = JSONObject.parseObject(content);
        JSONArray dataArray = jsonObject.getJSONArray("data");
        for (Object o : dataArray) {
            JSONObject x = (JSONObject) o;
            String name = x.getString("name");
            Long code = x.getLong("code");
            System.out.println("工作流程CODE为：" + code + "，工作流程名为：" + name);
        }
        response.close();
    }

    @Test
    public void testCreateProcessDefinition() throws IOException, ParseException {
        HttpPost request = new HttpPost(API_CONNECT_URL + "/projects/6576358231488/process-definition");
        request.setHeader(USER_TOKEN_KEY, USER_TOKEN);

        // set parameters
        List<NameValuePair> parameters = new ArrayList<>();
        parameters.add(new BasicNameValuePair("name", "pd2"));
        parameters.add(new BasicNameValuePair("locations", "[{\"taskCode\":6576405176384,\"x\":165,\"y\":88}]"));
        parameters.add(new BasicNameValuePair("taskDefinitionJson", "[{\"code\":6576405176384,\"delayTime\":\"0\",\"description\":\"\",\"environmentCode\":-1,\"failRetryInterval\":\"1\",\"failRetryTimes\":\"0\",\"flag\":\"YES\",\"name\":\"node_A\",\"taskParams\":{\"localParams\":[],\"rawScript\":\"echo \\\"test\\\"\",\"resourceList\":[]},\"taskPriority\":\"MEDIUM\",\"taskType\":\"SHELL\",\"timeout\":0,\"timeoutFlag\":\"CLOSE\",\"timeoutNotifyStrategy\":\"\",\"workerGroup\":\"default\"}]"));
        parameters.add(new BasicNameValuePair("taskRelationJson", "[{\"name\":\"\",\"preTaskCode\":0,\"preTaskVersion\":0,\"postTaskCode\":6576405176384,\"postTaskVersion\":0,\"conditionType\":\"NONE\",\"conditionParams\":{}}]"));
        parameters.add(new BasicNameValuePair("tenantCode", "default"));

        parameters.add(new BasicNameValuePair("description", "pd1"));
        parameters.add(new BasicNameValuePair("executionType", "PARALLEL"));
        parameters.add(new BasicNameValuePair("globalParams", "[]"));
        parameters.add(new BasicNameValuePair("timeout", "0"));
        UrlEncodedFormEntity formEntity = new UrlEncodedFormEntity(parameters);
        request.setEntity(formEntity);

        CloseableHttpResponse response = httpclient.execute(request);
        String content = EntityUtils.toString(response.getEntity(), "UTF-8");
        System.out.println(content);
        JSONObject jsonObject = JSONObject.parseObject(content);
        JSONObject data = jsonObject.getJSONObject("data");
        Long code = data.getLong("code");
        String name = data.getString("name");

        System.out.println("创建的工作流程名为：" + name);
        System.out.println("创建的工作流程CODE为：" + code);
        response.close();
    }




}
