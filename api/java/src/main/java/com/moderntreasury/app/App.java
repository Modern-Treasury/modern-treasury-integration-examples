/* Import Spark and MT client library */

package com.moderntreasury.examples;

import static spark.Spark.*;

import com.moderntreasury.api.client.ModernTreasuryClient;
import com.moderntreasury.api.client.okhttp.ModernTreasuryOkHttpClient;

import com.moderntreasury.api.models.*;

import io.github.cdimascio.dotenv.Dotenv;

import static spark.debug.DebugScreen.enableDebugScreen;

import java.util.List;
import java.util.Arrays;

public class App {
  @SuppressWarnings("deprecation")
  public static void main(String[] args) {
    /* Configure the Modern Treasury client with your API key and Org. ID */
    Dotenv dotenv = Dotenv.load();
    String apiKey = dotenv.get("MT_API_KEY");
    String mtOrgId = dotenv.get("MT_ORG_ID");
    String publicKey = dotenv.get("MT_PUB_KEY");
    String publicDir = dotenv.get("PUBLIC_DIR_PATH");

    enableDebugScreen();

    ModernTreasuryClient client = ModernTreasuryOkHttpClient.builder()
    .apiKey(apiKey)
    .organizationId(mtOrgId)
    .build();

    setPort(9001);
    if (publicDir != null) {
      staticFiles.externalLocation(publicDir);
    } else {
      externalStaticFileLocation("../../public");
    }

    /* POST route to handle a new counterpary and account collection form */
     post("/api/create-cp-acf", (req, res) -> {
      /* Create a counterpary */
      CounterpartyCreateParams cpParams = CounterpartyCreateParams.builder()
        .name(req.queryParams("name"))
        .build();
    
      Counterparty counterParty = client.counterparties().create(cpParams);

      List<String> rails = Arrays.asList(req.queryParamsValues("rails[]"));

      AccountCollectionFlowCreateParams acfParams = AccountCollectionFlowCreateParams.builder()
        .counterpartyId(counterParty.id())
        .paymentTypes(rails)
        .build();

      AccountCollectionFlow accountCollectionFlow = client.accountCollectionFlows().create(acfParams);

      req.session();
      req.session().attribute("client_token",accountCollectionFlow.clientToken().get());
      res.redirect("/embed.html");

      return res;
    });

    /* POST route to handle a new counterpary and payment form */
    post("/api/create-cp-pf", (req, res) -> {
      /* Create a counterpary */
      CounterpartyCreateParams cpParams = CounterpartyCreateParams.builder()
        .name(req.queryParams("name"))
        .build();

      Counterparty counterParty = client.counterparties().create(cpParams);

      PaymentFlowCreateParams.Direction direction = PaymentFlowCreateParams.Direction.of(req.queryParams("direction"));
      long amount = (long) (Double.parseDouble(req.queryParams("amount")) * 100);

      PaymentFlowCreateParams pfParams = PaymentFlowCreateParams.builder()
        .counterpartyId(counterParty.id())
        .amount(amount)
        .currency(req.queryParams("currency"))
        .direction(direction)
        .originatingAccountId(req.queryParams("originating_account_id"))
        .build();

      PaymentFlow paymentFlow = client.paymentFlows().create(pfParams);

      req.session();
      req.session().attribute("client_token",paymentFlow.clientToken().get());
      res.redirect("/embed.html");
      
      return res;
    });

    /* This endpoint provides configuration to modern-treasury-js */
    get("/config", (req, res) -> {
      res.type("application/javascript");
      res.body("window.mtConfig = { publishableKey: '"+ publicKey + "', clientToken: '" + req.session().attribute("client_token") + "' }");
      return res.body();
    });
  }
}