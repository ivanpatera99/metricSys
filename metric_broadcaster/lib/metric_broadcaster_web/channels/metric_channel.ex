defmodule MetricBroadcasterWeb.MetricChannel do
  use Phoenix.Channel

  def join("metric:lobby", _message, socket) do
    {:ok, %{welcome: "Welcome to the Metric Channel!"}, socket}
  end
end
