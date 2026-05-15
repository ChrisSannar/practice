namespace Async.Examples;

public class ConfigureAwaitExample
{
    public async Task ProcessAsync()
    {
        await Step1Async().ConfigureAwait(false);
        await Step2Async().ConfigureAwait(false);
    }

    public async Task ProcessWithCancellationAsync(CancellationToken token)
    {
        await Task.Delay(1000, token).ConfigureAwait(false);
        Console.WriteLine("Completed with cancellation support");
    }

    public async Task<int> CalculateWithTokenAsync(CancellationToken token)
    {
        await Task.Delay(50, token);
        return 42;
    }

    private async Task Step1Async()
    {
        await Task.Delay(50);
    }

    private async Task Step2Async()
    {
        await Task.Delay(50);
    }
}