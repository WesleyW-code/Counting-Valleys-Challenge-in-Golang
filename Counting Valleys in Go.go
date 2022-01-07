using System.CodeDom.Compiler;
using System.Collections.Generic;
using System.Collections;
using System.ComponentModel;
using System.Diagnostics.CodeAnalysis;
using System.Globalization;
using System.IO;
using System.Linq;
using System.Reflection;
using System.Runtime.Serialization;
using System.Text.RegularExpressions;
using System.Text;
using System;

class Result
{

    /*
     * Complete the 'countingValleys' function below.
     *
     * The function is expected to return an INTEGER.
     * The function accepts following parameters:
     *  1. INTEGER steps
     *  2. STRING path
     */

    public static int countingValleys(int steps, string path)
    {
        int valleys = 0;
        
        int alttitude = 0;
         
        List<int> alt_list = new List<int>();
        
        foreach (int value in Enumerable.Range(0, steps))
        {
            if (path[value] == 'U')
        {
            alttitude += 1;
        }
            else if(path[value] == 'D')
        {
            alttitude -= 1;
        }
        alt_list.Add(alttitude);
        }
        foreach (int value2 in Enumerable.Range(0, steps))
        {
            if (alt_list[value2] == 0 && alt_list [value2-1]<0)
        {
            valleys += 1;
        }
            
        }
             
    return valleys;
    }

}


class Solution
{
    public static void Main(string[] args)
    {
        TextWriter textWriter = new StreamWriter(@System.Environment.GetEnvironmentVariable("OUTPUT_PATH"), true);

        int steps = Convert.ToInt32(Console.ReadLine().Trim());

        string path = Console.ReadLine();

        int result = Result.countingValleys(steps, path);

        textWriter.WriteLine(result);

        textWriter.Flush();
        textWriter.Close();
    }
}
