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
            var parameters = new List<SqlParameter>
            {
                new SqlParameter("@documentType", "audit")
            };

            var query = new SqlQuerySpec("SELECT TOP 1 c.RunID, c._ts FROM c WHERE c.DocumentType= @documentType ORDER BY c._ts DESC",
                            new SqlParameterCollection(parameters));

            var result = Client.CreateDocumentQuery<Run>(
                ReportsLink,
                query, 
                FeedOptions)
                .AsEnumerable();
            return result.First();
        }

        public IList<Audit> GetReports(string runID)
        {
            var parameters = new List<SqlParameter>
            {
                new SqlParameter("@runID", runID),
                new SqlParameter("@documentType", "audit")
            };

            var query = new SqlQuerySpec("SELECT * FROM c WHERE c.RunID = @runID AND c.DocumentType = @documentType",
                    new SqlParameterCollection(parameters));

            var result = Client.CreateDocumentQuery<Audit>(ReportsLink, query, FeedOptions);
            return result.ToList();
        }
    }
}
