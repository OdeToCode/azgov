using System;
using Microsoft.Azure.Documents;
using Newtonsoft.Json;

namespace GovenorReports.Data
{
    public class Run
    {
        public string RunID { get; set; }

        [JsonPropertyAttribute(PropertyName = "_ts")]
        [JsonConverterAttribute(typeof(UnixDateTimeConverter))]
        public DateTime Runtime { get; set; }
    }
}
