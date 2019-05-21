using System.Collections.Generic;
public interface Product {
    Product clone();
    int get_cost();
}

public class Locomotive : Product {
    private int max_speed;
    private int price;
    public string color;
    public Locomotive(int max_speed) {
        this.max_speed = max_speed;
        price = 15;
        color = "red";
    }

    public void set_price(int price){
        this.price = price;
    }

    public void set_color(string color) {
        this.color = color;
    }

    public int get_cost() {
        return price;
    }

    public Product clone() {
        return (Product) this.MemberwiseClone();
    }
}

public class Train : Product {
    private Locomotive locomotive;
    private List<string> wagons = new List<string>();
    public Train(Locomotive locomotive) {
        this.locomotive = locomotive;
    }

    public void add_wagon(string wagon) {
        wagons.Add(wagon);
    }

    public Product clone() {
        Train clone = (Train) this.MemberwiseClone();
        clone.wagons = new List<string>(wagons);
        clone.locomotive = (Locomotive) locomotive.clone();
        return clone;
    }

    public int get_cost() {
        return locomotive.get_cost() + 5*wagons.Count;
    }
}

public class Tester {
    public static void Main(string[] args) {
        Locomotive oldLocomotive = new Locomotive(15);
        Product newLocomotive = oldLocomotive.clone();
        Train longTrain = new Train(oldLocomotive);
        longTrain.add_wagon("Passenger");
        Product shortTrain = longTrain.clone();
        longTrain.add_wagon("Freight");
        
        oldLocomotive.set_price(7);

        System.Console.WriteLine(newLocomotive.get_cost() - oldLocomotive.get_cost());
        System.Console.WriteLine(shortTrain.get_cost() - longTrain.get_cost());
    }
}