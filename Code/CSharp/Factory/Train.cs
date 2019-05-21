using System;

abstract class Train {
    private int id;
    private string color;

    public Train(int id, string color){
        this.id = id;
        this.color = color;
    }

    public void describe() {
        Console.WriteLine("I am a train with id = " + id + " and color = " + color);
    }
}