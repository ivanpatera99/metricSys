defmodule MetricBroadcasterWeb.UserSocket do
  use Phoenix.Socket

  ## Channels
  channel "metric:lobby", MetricBroadcasterWeb.MetricChannel

   ## Transports
   transport :websocket, Phoenix.Transports.WebSocket
   transport :longpoll, Phoenix.Transports.LongPoll

  def connect(_params, socket, _connect_info) do
    {:ok, socket}
  end
  def id(_socket), do: nil
end
