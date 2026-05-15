using System;
using System.Collections.Generic;
using System.Linq;

var products = new List<Product>
{
    new(1, "Laptop", 999.99m, "Electronics"),
    new(2, "Mouse", 29.99m, "Electronics"),
    new(3, "Chair", 149.99m, "Furniture"),
    new(4, "Desk", 299.99m, "Furniture"),
    new(5, "Monitor", 399.99m, "Electronics"),
    new(6, "Keyboard", 79.99m, "Electronics")
};

var users = new List<User>
{
    new(1, "Alice", "alice@example.com"),
    new(2, "Bob", "bob@example.com"),
    new(3, "Charlie", "charlie@example.com"),
    new(4, "Alice", "alice2@example.com")
};

var orders = new List<Order>
{
    new(1, 1, 999.99m),
    new(2, 1, 29.99m),
    new(3, 2, 149.99m),
    new(4, 3, 299.99m)
};

Console.WriteLine("=== LINQ Demo ===\n");

Console.WriteLine("--- Select (Projection) ---");
var names = products.Select(p => p.Name);
Console.WriteLine($"Product names: {string.Join(", ", names)}");

Console.WriteLine("\n--- Where (Filtering) ---");
var expensiveProducts = products.Where(p => p.Price > 100);
Console.WriteLine($"Expensive products: {string.Join(", ", expensiveProducts.Select(p => p.Name))}");

Console.WriteLine("\n--- OrderBy / OrderByDescending ---");
var orderedByPrice = products.OrderBy(p => p.Price).Select(p => $"{p.Name} (${p.Price})");
Console.WriteLine($"Ordered by price: {string.Join(", ", orderedByPrice)}");

var orderedDesc = products.OrderByDescending(p => p.Price).Select(p => p.Name);
Console.WriteLine($"Ordered descending: {string.Join(", ", orderedDesc)}");

Console.WriteLine("\n--- GroupBy ---");
var groupedByCategory = products.GroupBy(p => p.Category);
foreach (var group in groupedByCategory)
{
    Console.WriteLine($"  {group.Key}: {string.Join(", ", group.Select(p => p.Name))}");
}

Console.WriteLine("\n--- Join ---");
var orderDetails = orders.Join(products, o => o.ProductId, p => p.Id, (o, p) => new { o.OrderId, Product = p.Name, Price = o.Amount });
foreach (var detail in orderDetails)
{
    Console.WriteLine($"  Order {detail.OrderId}: {detail.Product} - ${detail.Price}");
}

Console.WriteLine("\n--- Aggregate (Sum, Average, Count, Max, Min) ---");
Console.WriteLine($"Total price: ${products.Sum(p => p.Price):F2}");
Console.WriteLine($"Average price: ${products.Average(p => p.Price):F2}");
Console.WriteLine($"Count: {products.Count}");
Console.WriteLine($"Max price: ${products.Max(p => p.Price):F2}");
Console.WriteLine($"Min price: ${products.Min(p => p.Price):F2}");

Console.WriteLine("\n--- First / FirstOrDefault / Single / SingleOrDefault ---");
var first = products.First(p => p.Category == "Electronics");
Console.WriteLine($"First electronics: {first.Name}");
var notFound = products.FirstOrDefault(p => p.Price > 10000);
Console.WriteLine($"FirstOrDefault (not found): {notFound?.Name ?? "null"}");

Console.WriteLine("\n--- Any / All / Contains ---");
Console.WriteLine($"Any expensive (>500): {products.Any(p => p.Price > 500)}");
Console.WriteLine($"All affordable (<2000): {products.All(p => p.Price < 2000)}");
Console.WriteLine($"Contains 'Mouse': {products.Select(p => p.Name).Contains("Mouse")}");

Console.WriteLine("\n--- Skip / Take (Pagination) ---");
var page1 = products.Skip(0).Take(3);
Console.WriteLine($"Page 1: {string.Join(", ", page1.Select(p => p.Name))}");
var page2 = products.Skip(3).Take(3);
Console.WriteLine($"Page 2: {string.Join(", ", page2.Select(p => p.Name))}");

Console.WriteLine("\n--- Distinct ---");
var uniqueNames = users.Select(u => u.Name).Distinct();
Console.WriteLine($"Unique user names: {string.Join(", ", uniqueNames)}");

Console.WriteLine("\n--- SelectMany ---");
var allChars = products.SelectMany(p => p.Name.ToLower().Distinct()).Distinct().OrderBy(c => c);
Console.WriteLine($"All unique chars in product names: {string.Join("", allChars)}");

Console.WriteLine("\n--- Let (query syntax) ---");
var discounted = from p in products
                 let discount = p.Price * 0.9m
                 where discount > 100
                 select new { p.Name, OriginalPrice = p.Price, DiscountedPrice = discount };
foreach (var item in discounted)
{
    Console.WriteLine($"  {item.Name}: ${item.OriginalPrice} -> ${item.DiscountedPrice:F2}");
}

Console.WriteLine("\n=== Demo Complete ===");

Console.WriteLine("\n=== Running Tests ===\n");

int passed = 0;
int failed = 0;

void RunTest(string name, Action test)
{
    try
    {
        test();
        Console.WriteLine($"[PASS] {name}");
        passed++;
    }
    catch (Exception ex)
    {
        Console.WriteLine($"[FAIL] {name}: {ex.Message}");
        failed++;
    }
}

RunTest("Select_ProjectsToNewType", () =>
{
    var result = products.Select(p => p.Name).ToList();
    Assert.Equal(6, result.Count);
    Assert.Contains("Laptop", result);
});

RunTest("Where_FiltersCorrectly", () =>
{
    var result = products.Where(p => p.Category == "Electronics").ToList();
    Assert.Equal(4, result.Count);
});

RunTest("OrderBy_SortsAscending", () =>
{
    var result = products.OrderBy(p => p.Price).First();
    Assert.Equal("Mouse", result.Name);
});

RunTest("GroupBy_GroupsByCategory", () =>
{
    var result = products.GroupBy(p => p.Category).ToList();
    Assert.Equal(2, result.Count);
    Assert.Equal(4, result.First(g => g.Key == "Electronics").Count());
});

RunTest("Join_CombinesData", () =>
{
    var result = orders.Join(products, o => o.ProductId, p => p.Id, (o, p) => p.Name).ToList();
    Assert.Contains("Laptop", result);
});

RunTest("Aggregate_Sum", () =>
{
    var total = products.Sum(p => p.Price);
    Assert.Equal(1959.94m, total);
});

RunTest("Any_ReturnsBoolean", () =>
{
    var hasExpensive = products.Any(p => p.Price > 500);
    Assert.True(hasExpensive);
});

RunTest("All_ReturnsBoolean", () =>
{
    var allAffordable = products.All(p => p.Price < 2000);
    Assert.True(allAffordable);
});

RunTest("SkipTake_Paginates", () =>
{
    var page1 = products.Skip(0).Take(3).ToList();
    Assert.Equal(3, page1.Count);
    Assert.Equal("Laptop", page1[0].Name);
});

Console.WriteLine($"\n=== Test Results: {passed} passed, {failed} failed ===");

static class Assert
{
    public static void Equal<T>(T expected, T actual) 
    {
        if (!EqualityComparer<T>.Default.Equals(expected, actual))
            throw new Exception($"Expected {expected}, got {actual}");
    }
    public static void True(bool condition) 
    {
        if (!condition) throw new Exception("Expected true");
    }
    public static void Contains<T>(T item, IEnumerable<T> collection) 
    {
        if (!collection.Contains(item))
            throw new Exception($"Collection should contain {item}");
    }
}

record Product(int Id, string Name, decimal Price, string Category);
record User(int Id, string Name, string Email);
record Order(int OrderId, int ProductId, decimal Amount);