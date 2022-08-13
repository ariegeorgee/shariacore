package ariegeorgee.shariacore;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import ariegeorgee.shariacore.service.CoreService;

@SpringBootApplication
public class ShariacoreApplication {

	public static void main(String[] args) {
		SpringApplication.run(ShariacoreApplication.class, args);

		CoreService service = new CoreService();
		service.processRequest();
		System.exit(0);
	}

}
