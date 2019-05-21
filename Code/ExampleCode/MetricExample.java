//Example code used for examplifying the application of the metrics used in the thesis.

interface Animal {
    public void speak();
}

abstract class Canine implements Animal {
    protected String sound;
    public void speak() {
        System.out.println(this.sound);
    }
}

class Dog extends Canine {
    public Dog() {
        sound = "Woof";
    }
}

class Wolf extends Canine {
    public Wolf() {
        sound = "Howl";
    }
}

class Pig implements Animal {
    private String oldSound = "Oink";
    private String youngSound = "Squee";
    public int age;

    public Pig(int age) {
        this.age = age;
    }

    public void speak() {
        if (age < 2) {
            System.out.println(youngSound);
        }
        else {
            System.out.println(oldSound);
        }
    }
}

class MetricExample {
    public static void main(String[] args) {
        Animal[] animals = new Animal[]{new Dog(), new Wolf(), new Pig(1), new Pig(2)};
        for (Animal animal : animals) {
            animal.speak();
        } 
    }
}