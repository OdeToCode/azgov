using System.Collections.Generic;
using Newtonsoft.Json;
using Newtonsoft.Json.Linq;

namespace GovenorReports.Data
{
    public class AuditReport
    {
        public string ResourceID { get; set; }
        public string SubscriptionID { get; set; }
        public string GroupName { get; set; }
        public string Name { get; set; }
        public string Type { get; set; }
        public string RunID { get; set; }
        public bool Failed { get; set; }

        [JsonExtensionData]
        public IDictionary<string, JToken> Properties;
    }
}