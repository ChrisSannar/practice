namespace DependencyInjection.Services;

public interface IMessageService
{
    string GetMessage();
    void SendMessage(string message);
}