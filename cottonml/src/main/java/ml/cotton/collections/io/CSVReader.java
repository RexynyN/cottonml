package src.ml.cotton.collections.io;

import java.io.*;
import java.nio.file.*;
import java.util.ArrayList;
import java.util.List;

public class CSVReader {
    private Path filePath; 

    public List<String> readRawCSV() {
        List<String> records = null; 

        try {
            int listSize = (int) Files.lines(this.filePath).count();
            records = new ArrayList<String>(listSize);

            BufferedReader reader = Files.newBufferedReader(filePath);
        } catch(IOException e) {
            System.out.println("ERROR: " + this.filePath + " is not a valid file or doesn't exists!");
        }
        

        return records;
    }

    public Path getFilePath() {
        return filePath;
    }

    public void setFilePath(Path filePath) {
        this.filePath = filePath;
    }
    
    public void setFilePath(String filePath) {
        this.filePath = Paths.get(filePath);
    }
}

