# bqsql

**bqsql** is a command-line tool for easily running SQL queries against Google BigQuery tables. It simplifies the process of querying and managing BigQuery data by allowing you to write SQL queries directly in your terminal.

## Installation

To install **bqsql**, you need to have Go installed on your machine. You can download and install Go from the [official website](https://golang.org/dl/).

Once you have Go installed, you can install **bqsql** using `go get`:

```bash
go get github.com/rantolin/bqsql
```

This command will fetch the source code, build the executable, and place it in your Go bin directory. Make sure your Go bin directory is included in your system's PATH so you can run **bqsql** from anywhere in your terminal.

Alternatively, you can clone the repository and build the executable manually:

```bash
git clone https://github.com/rantolin/bqsql.git
cd bqsql
go build
```

This will create an executable named `bqsql` in the current directory. You can move this executable to a directory in your PATH to make it accessible from anywhere in your terminal.

After installation, you can verify that **bqsql** is installed correctly by running:

```bash
bqsql --version
```

This should display the version of **bqsql** installed on your system.

## Usage

### Authentication

Before using **bqsql**, make sure you have authenticated your Google Cloud account and have the necessary permissions to access your BigQuery datasets and tables. You can authenticate by running:

```bash
gcloud auth application-default login
```

This command will open a browser window for you to authenticate.

### Running Queries

To run a SQL query against a BigQuery table, use the following command:

```bash
bqsql query "<SQL_QUERY>"
```

For example:

```bash
bqsql query 'SELECT * FROM `project_id.dataset.my_table`;'
```

Replace `<SQL_QUERY>` with your SQL query enclosed in double or single quotes.

However, if you want to get a flavor of what the table looks like you can use the `head` command to query 10 records:

```bash
bqsl --project_id <PROJECT_ID> --dataset <DATASET_ID> head "my_table"
```

Replace `<PROJECT_ID>` with your Google Cloud project ID, `<DATASET_ID>` with the BigQuery dataset ID containing the table you want to query.

### Listing Resources
The list command allows you to list resources in your BigQuery project. Currently, it supports listing datasets and tables.

To list all datasets in a project, use the following command:

```bash
bqsql --project_id <PROJECT_ID> list datasets
```

Replace <PROJECT_ID> with your Google Cloud project ID. This command will return a list of all datasets in the specified project.

To list all tables in a dataset, use the following command:
```bash
bqsql --project_id <PROJECT_ID> --dataset <DATASET_ID> list tables
```

Replace <PROJECT_ID> with your Google Cloud project ID and <DATASET_ID> with the BigQuery dataset ID for which you want to list tables. This command will return a list of all tables in the specified dataset.

### Help

To see all available options and commands, you can use the `--help` flag:

```bash
bqsql --help
```

You can use the `--help` flag to get aditional options and commands of subcommands:

```bash
bqsql describe --help
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Contributing

Contributions are welcome! Please feel free to open an issue or submit a pull request.

---

Feel free to modify or expand upon this README as needed!
