namespace DependencyInjection.Services;

public class MessageService : IMessageService
{
    private readonly string _prefix;

    public MessageService()
    {
        _prefix = "Default";
    }

    public MessageService(string prefix)
    {
        _prefix = prefix;
    }

    public string GetMessage() => $"{_prefix}: Hello from MessageService!";

    public void SendMessage(string message) => Console.WriteLine($"[MessageService] {message}");
}