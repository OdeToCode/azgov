using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;
using GovenorReports.Data;
using Microsoft.AspNetCore.Mvc;
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

        public void OnGet()
        {
            Run = queries.GetLastRun();
            Costs = queries.GetCostReports(Run.RunID);
        }
    }
}