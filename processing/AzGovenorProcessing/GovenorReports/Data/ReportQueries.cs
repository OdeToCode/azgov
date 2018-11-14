using System;
using System.Collections.Generic;
using System.Linq;
using Microsoft.Azure.Documents;
using Microsoft.Azure.Documents.Client;

namespace GovenorReports.Data
{
    public class ReportQueries
    {
        public ReportQueries(DocumentClient client)
        {
            Client = client;
            ReportsLink = UriFactory.CreateDocumentCollectionUri("govenor", "reports");
            FeedOptions = new FeedOptions { EnableCrossPartitionQuery = true };
        }

        public DocumentClient Client { get; }
        public Uri ReportsLink { get; set; }
        public FeedOptions FeedOptions { get; set; }

        public Run GetLastRun()
        {
            var query = Client.CreateDocumentQuery<Run>(
                ReportsLink,
                "SELECT TOP 1 c.RunID, c._ts FROM c ORDER BY c._ts DESC", 
                FeedOptions)
                .AsEnumerable();
            return query.First();
        }

        public IList<Report> GetReports(string runID)
        {
            var parameters = new List<SqlParameter>
            {
                new SqlParameter("@runID", runID)
            };

            var query = new SqlQuerySpec("SELECT * FROM c WHERE c.RunID = @runID",
                    new SqlParameterCollection(parameters));

            var result = Client.CreateDocumentQuery<Report>(ReportsLink, query, FeedOptions);
            return result.ToList();
        }
    }
}
