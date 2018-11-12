using Microsoft.Azure.WebJobs;
using Microsoft.Extensions.Logging;

namespace EventSinks
{
    public static class ReportSink
    {
        [FunctionName("ProcessReport")]
        public static void Run([EventHubTrigger("azgovhub", Connection = "GovHub")]
            string myEventHubMessage, ILogger log)
        {
            log.LogInformation($"C# Event Hub trigger function processed a message: {myEventHubMessage}");
        }
    }
}
