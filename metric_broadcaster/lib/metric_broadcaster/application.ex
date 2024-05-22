defmodule MetricBroadcaster.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      MetricBroadcasterWeb.Telemetry,
      {DNSCluster, query: Application.get_env(:metric_broadcaster, :dns_cluster_query) || :ignore},
      {Phoenix.PubSub, name: MetricBroadcaster.PubSub},
      # Start a worker by calling: MetricBroadcaster.Worker.start_link(arg)
      # {MetricBroadcaster.Worker, arg},
      # Start to serve requests, typically the last entry
      MetricBroadcasterWeb.Endpoint,
      MetricBroadcasterWeb.MetricFetcher
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: MetricBroadcaster.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    MetricBroadcasterWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
