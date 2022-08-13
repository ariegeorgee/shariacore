package ariegeorgee.shariacore.service;

import ariegeorgee.shariacore.dao.BankData;
import ariegeorgee.shariacore.dao.BankDataDao;
import java.util.ArrayList;
import java.util.LinkedHashSet;
import java.util.List;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.ExecutorService;
import java.util.concurrent.Executors;

public class CoreService {

    public void processRequest() {
        BankDataDao dao = new BankDataDao();
        List<BankData> data;
        try {
            data = dao.getBankData();
        } catch (Exception e) {
            //Add monitoring alert
            System.out.println("Failed to read data " + e);
            return;
        }

        final ExecutorService execAsync1 = Executors.newFixedThreadPool(20);
        final ExecutorService execAsync2 = Executors.newFixedThreadPool(8);
        CountDownLatch waitGroup1 = new CountDownLatch(data.size());

        //Handle Question 1, 2a, and 2b
        for (BankData row : data) {
            execAsync1.execute(() -> {
                int threadNum = (int) Thread.currentThread().getId();
                soal1Action(row, threadNum);
                soal2Action(row, threadNum);
                waitGroup1.countDown();
            });
        }

        try {
            waitGroup1.await();
        } catch (InterruptedException e) {
            System.out.println("Failed to process on question 1 to 2b " + e);
            //Add monitoring alert
        }

        //Handle question 3
        CountDownLatch waitGroup2 = new CountDownLatch(data.size());
        for (BankData row : data) {
            final int additional = data.indexOf(row) < 100 ? 10 : 0;
            execAsync2.execute(() -> {
                int threadNum = (int) Thread.currentThread().getId();
                soal3Action(row, threadNum, additional);
                waitGroup2.countDown();
            });
        }

        try {
            waitGroup2.await();
        } catch (InterruptedException e) {
            System.out.println("Failed to process on question 3 " + e);
        }
        dao.saveBankData(data);
    }

    private void soal1Action(BankData data, int threadNum) {
        data.setNo1(threadNum);
        float avg = (data.getBalance() + data.getPreviousBalance()) / 2f;
        data.setAverageBalance(avg);
    }

    private void soal2Action(BankData data, int threadNum) {
        data.setNo2a(threadNum);
        data.setNo2b(threadNum);

        if (data.getBalance() >= 100 && data.getBalance() <= 150) {
            data.setFreeTransfer(5);
        }
        if (data.getBalance() > 150) {
            data.setBalance(data.getBalance() + 25);
        }
    }

    private void soal3Action(BankData data, int threadNum, int additional) {
        data.setNo3(threadNum);
        data.setBalance(data.getBalance() + additional);
    }
}
