using System.Collections.Generic;
using System.Linq;
using GovenorReports.Data;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace GovenorReports.Pages
{
    public class IndexModel : PageModel
    {
        private readonly ReportQueries queries;

        public Run LastRun { get; private set; }
        public IList<Audit> Reports { get; private set; }
        public IEnumerable<string> Subscriptions { get; set; }
        public IEnumerable<Audit> WebApps { get; private set; }
        public IEnumerable<Audit> SqlServers { get; private set; }
        public IEnumerable<Audit> RedisCaches { get; private set; }

        const string RedisType = "Microsoft.Cache/Redis";
        const string SqlType = "Microsoft.Sql/servers";
        const string AppService = "Microsoft.Web/sites";


        public IndexModel(ReportQueries queries)
        {
            this.queries = queries;
        }

        public void OnGet()
        {
            LastRun = queries.GetLastRun();
            Reports = queries.GetReports(LastRun.RunID);
            Subscriptions = Reports.Select(r => r.SubscriptionID).Distinct();
            WebApps = Reports.Where(r => r.Type == AppService);
            SqlServers = Reports.Where(r => r.Type == SqlType);
            RedisCaches = Reports.Where(r => r.Type == RedisType);
        }
    }
}
