defmodule MetricBroadcasterWeb.Router do
  use MetricBroadcasterWeb, :router

  pipeline :api do
    plug :accepts, ["json"]
  end

  scope "/api", MetricBroadcasterWeb do
    pipe_through :api
  end
end
