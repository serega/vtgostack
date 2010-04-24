
package vtgostack

import "testing"


func TestArrayStack(t *testing.T) {
    var arrayStack Stack = NewArrayStack(3)
    testImpl(t, arrayStack, 3)
}

func TestLinkedListStack(t *testing.T) {
    linkedListStack := NewLinkedListStack(3)
    testImpl(t, linkedListStack, 3)
}

func assertValue(t *testing.T, expected int, actual int, message string) {
    if expected != actual {
        t.Errorf("expected %s, actual %s   %s", expected, actual, message)
    }
}

func testImpl(t *testing.T, s Stack, maxDepth int) {
    if s.Depth() != 0 {
        t.Errorf("Initial stack has depth %d", s.Depth())
        return
    }
    assertValue(t, maxDepth, s.RemCapacity(), "RemCapacity()")
    s.Push(1)
    assertValue(t, 1, s.Top().(int), " Top()")
    assertValue(t, 1, s.Depth(), "Depth()")
    assertValue(t, 2, s.RemCapacity(), "RemCapacity()")
    
    s.Push(3)
    assertValue(t, 3, s.Top().(int), " Top()")
    assertValue(t, 2, s.Depth(), "Depth()")
    assertValue(t, 1, s.RemCapacity(), "RemCapacity()")
    
    s.Push(5)
    assertValue(t, 3, s.Depth(), "Depth()")
    assertValue(t, 0, s.RemCapacity(), "RemCapacity()")
    
    s.Push(6) // does nothing
    assertValue(t, 5, s.Top().(int), " Top()")
    assertValue(t, 5, s.Pop().(int), " Pop()")
    assertValue(t, 2, s.Depth(), "Depth()")
    assertValue(t, 1, s.RemCapacity(), "RemCapacity()")
    
    
    assertValue(t, 3, s.Pop().(int), " Pop()")
    assertValue(t, 1, s.Depth(), "Depth()")
    assertValue(t, 2, s.RemCapacity(), "RemCapacity()")
    
    assertValue(t, 1, s.Pop().(int), " Pop()")
    assertValue(t, 0, s.Depth(), "Depth()")
    assertValue(t, 3, s.RemCapacity(), "RemCapacity()")    
}


func TestFlip(t *testing.T) {
    var flippingStack FlippingStack = &FlippingArrayStack{ NewArrayStack(3) }
    testFlip(t, flippingStack)
    
    flippingStack = &FlippingLinkedListStack{ NewLinkedListStack(3) }
    testFlip(t, flippingStack)
}


func testFlip(t *testing.T, flipping FlippingStack) {
    flipping.Push(1)
    flipping.Push(2)
    flipping.Push(3)
    
    flipped := flipping.Flip()
    assertValue(t, 1, flipped.Pop().(int), "Pop()")
    assertValue(t, 2, flipped.Pop().(int), "Pop()")
    assertValue(t, 3, flipped.Pop().(int), "Pop()")
    assertValue(t, 0, flipped.Depth(), "Depth()")
}



