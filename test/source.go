func Integer: numPrint (Integer: num, Integer: length)
{
    Integer: i, j, first, temp;
    char : a;
    a := 'x';
    print ( "enter number" );
    In>> i;
    println (i);
    i := length;
    while i > 0:
    {
        first:=0;
        j <- 1;
        while j < i:
        {
            write( j);
            j:=j + 1;
        }
        if j == 1:{
            print("one");
        }
        elif j==2:{
            print("two");
        }
        else
        {
            print("others");
            }

        /* this is a comment */
        i:=i - 1;
        /*This is a
                Multiline
                Comment*/
        }
        print( "temp is");
        println(temp);
        ret i;
}