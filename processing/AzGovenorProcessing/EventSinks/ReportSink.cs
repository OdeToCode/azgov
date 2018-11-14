using System.Threading.Tasks;
using Microsoft.Azure.WebJobs;
using Microsoft.Extensions.Logging;
using Newtonsoft.Json.Linq;

namespace EventSinks
{
    public static class ReportSink
    {
        [FunctionName("ProcessReport")]
        public static async Task Run(
            [EventHubTrigger(eventHubName: "azgovhub", Connection = "GovHub")]
            string myEventHubMessage, 
            [CosmosDB(databaseName:"govenor", collectionName:"reports", ConnectionStringSetting ="AzGovenorDb")]
            IAsyncCollector<JObject> output,
            ILogger log)
        {
            var data = JObject.Parse(myEventHubMessage);
            await output.AddAsync(data);
            log.LogInformation($"C# Event Hub trigger function processed a message: {myEventHubMessage}");
        }
    }
}
