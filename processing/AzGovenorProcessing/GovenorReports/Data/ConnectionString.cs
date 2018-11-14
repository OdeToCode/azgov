using System;

namespace GovenorReports.Data
{
    public class ConnectionString
    {
        public ConnectionString(string connectionString)
        {
            if (string.IsNullOrEmpty(connectionString))
            {
                throw new ArgumentException(nameof(connectionString));
            }

            var tokens = connectionString.Split(";");
            if (tokens.Length != 3)
            {
                throw new ArgumentException($"Could not parse URI and Key from {nameof(connectionString)}");
            }

            ServiceUri = new Uri(tokens[0]);
            AuthKey = tokens[1].Replace("AccountKey=", "");
        }

        public Uri ServiceUri { get; protected set; }
        public string AuthKey { get; protected set; }
    }
}
