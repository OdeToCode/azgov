using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace GovenorReports.Data
{
    public class CostReport
    {
        public string ResourceId { get; set; }
        public float Cost { get; set; }
        public float MonthlyCost {
            get
            {
                return Cost * (31 / 7);
            }
        }
    }
}
