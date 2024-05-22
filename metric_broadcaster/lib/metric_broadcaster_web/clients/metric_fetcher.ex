defmodule MetricBroadcasterWeb.MetricFetcher do
  use GenServer

  def start_link(_) do
    GenServer.start_link(__MODULE__, %{})
  end

  def init(_) do
    schedule_work()  # Schedule work to be done 1 second later
    {:ok, %{}}
  end

  def handle_info(:work, state) do
    # metrics = fetch_metrics()
    metrics = fetch_not_fetch()
    IO.puts("Broadcasting metrics: #{inspect(metrics)}")
    MetricBroadcasterWeb.Endpoint.broadcast("metric:lobby", "metrics_data", %{metrics: metrics})
    schedule_work()  # Reschedule the work
    {:noreply, state}
  end

  defp schedule_work do
    Process.send_after(self(), :work, 1000)  # Schedule work to be done 1 second later
  end

  defp fetch_not_fetch do
    %{
      cpu_usage: "#{Enum.random(1..100)}%",
      memory_usage: "#{Enum.random(1..100)}%",
      disk_io: "#{Enum.random(1..100)}%"
    }
  end

  defp fetch_metrics do
    # Perform an HTTP GET request to fetch metrics from a service running on port 8080
    url = "http://metrics-generator:8080/metrics?limit=1"
    IO.puts(("Fetching metrics from #{url}..."))
    response = HTTPoison.get!(url)
    case response.status_code do
      200 ->
        # Parse the response body and extract the metrics
        body = response.body
        metrics = parse_metrics(body)
        metrics
      _ ->
        # Handle error cases
        %{error: "Failed to fetch metrics"}
    end
  end

  defp parse_metrics(body) do
    case Jason.decode(body) do
      {:ok, %{"metrics" => metrics}} ->
        metrics
      {:error, _} ->
        %{error: "Failed to parse metrics"}
    end
  end
end
