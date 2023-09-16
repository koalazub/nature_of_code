# GnG(Graphics and Graphics)

In my quest to try and intersect everything I'm interested in - this project aims to refine my skills in maths and
programming so that I don't have to discretely learn each field with the hopes of freeing up some time(fat chance).

I like Go and want to get better at it. I also like maths and want to get better at it. So, in the spirit of closing
the  gap between my sanity and a mental ward - I've decided to tackle the two simulataneously with the hopes of getting
closer to grokking the two.

# Why tho?

Most of the jobs I'm enamoured by deal with some deep thinking. My critical thinking processes are all over the place.
If you've had the misfortune of coming across my [linkedin](https://www.linkedin.com/in/cokoali/) - you can see that I
have worked across myriad fields. This is my quest of maintaining focus and finally becoming a [master of one](https://
www.shaunnestor.com/jack-of-all-trades-quote/)

# Installation

Using `nix`? cheeky `nix develop` should be all you need

Otherwise, check the `nativeBuildInputs` set in my `flake.nix` file for the packages that are required. 
For the lazy: 
```nix 
  nativeBuildInputs = with pkgs; [
    gopls
    go_1_21
    vscode-langservers-extracted 
  ];  
```


