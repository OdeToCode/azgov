using System.Collections.Generic;
using System.Linq;
using GovenorReports.Data;
using Microsoft.AspNetCore.Mvc.RazorPages;

namespace GovenorReports.Pages
{
    public class CostModel : PageModel
    {
        private readonly ReportQueries queries;

        public CostModel(ReportQueries queries)
        {
            this.queries = queries;
        }

        public Run Run { get; private set; }
        public IList<CostReport> Costs { get; private set; }
        public IEnumerable<IGrouping<string, CostReport>> CostsByGroup { get; private set; }

        public void OnGet()
        {
            Run = queries.GetLastRun();
            Costs = queries.GetCostReports(Run.RunID);
            CostsByGroup = Costs.GroupBy(c => c.GroupName.ToLower()).OrderByDescending(g => g.Sum(r => r.MonthlyCost));
        }
    }
}