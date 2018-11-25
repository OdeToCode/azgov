using System;
using Microsoft.Azure.Documents.Client;
using Microsoft.Azure.WebJobs;
using Microsoft.Azure.WebJobs.Host;
using Microsoft.Extensions.Logging;

namespace EventSinks
{
    public static class CleanOldReports
    {
        [FunctionName("CleanOldReports")]
        public static void Run(
            [TimerTrigger("0 0 0 1 * *")]TimerInfo myTimer,
            [CosmosDB(databaseName: "govenor", collectionName: "reports", ConnectionStringSetting = "AzGovenorDb")]
            DocumentClient client,
            ILogger log)
        {
            log.LogInformation($"Cleaning out those old reports at: {DateTime.Now}");

            // TODO
        }
    }
}
