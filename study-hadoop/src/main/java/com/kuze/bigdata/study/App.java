package com.kuze.bigdata.study;

import org.apache.hadoop.conf.Configuration;
import org.apache.hadoop.fs.FSDataOutputStream;
import org.apache.hadoop.fs.FileSystem;
import org.apache.hadoop.fs.Path;

import java.io.IOException;

/**
 * Hello world!
 *
 */
public class App 
{
    public static void main( String[] args ) throws IOException, InterruptedException {
        Configuration conf = new Configuration();
        conf.set("fs.alluxio.impl", "alluxio.hadoop.FileSystem");
        conf.set("fs.AbstractFileSystem.alluxio.impl", "alluxio.hadoop.AlluxioFileSystem");
        conf.set("fs.defaultFS", "alluxio://localhost:19998");

        Path path = new Path("/test");

        Thread t1 = new Thread(new Runnable() {
            @Override
            public void run() {
                FileSystem fs = null;
                FSDataOutputStream output = null;
                try {
                    fs = FileSystem.newInstance(conf);
                    output = fs.append(path);
                    Thread.sleep(5000);
                    output.writeUTF("test1");
                    output.close();
                    fs.close();
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        });
        t1.start();

        Thread t2 = new Thread(new Runnable() {
            @Override
            public void run() {
                FileSystem fs = null;
                FSDataOutputStream output = null;
                try {
                    fs = FileSystem.newInstance(conf);
                    output = fs.append(path);
                    output.writeUTF("test2");
                    output.close();
                    fs.close();
                } catch (Exception e) {
                    e.printStackTrace();
                }
            }
        });
        Thread.sleep(1000);
        t2.start();
    }
}
