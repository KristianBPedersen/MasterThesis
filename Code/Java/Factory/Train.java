abstract class Train {
    private int id;
    private String color;

    public Train(int id, String color) {
        this.id = id;
        this.color = color;
    }

    public void describe() {
        System.out.println("I am a train with id = " + id + " and color = " + color);
    }
}