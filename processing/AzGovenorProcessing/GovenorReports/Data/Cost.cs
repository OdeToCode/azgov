using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GovenorReports.Data
{
    public class CostReport
    {
        public string ResourceID { get; set; }
        public string SubscriptionID { get; set; }
        public string GroupName { get; set; }
        public string Name { get; set; }
        public string Type { get; set; }
        public string RunID { get; set; }
        public float Cost { get; set; }
        public float MonthlyCost {
            get
            {
                return Cost * (31 / 7);
            }
        }
    }
}
