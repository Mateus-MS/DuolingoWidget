## Usage

Example:
```bash
<img src="https://cubonauta.com?id=your_user_id&mood=angry" width=150>
```

#### parameters

| Parameters   | Type       | Description                                   | Mandatory |
| :---------- | :--------- | :------------------------------------------ |:-------- |
| `id`        | `string`   | Your user ID [see how to get](#UserID). | `true`
| `mood`      | `string`   | The mood you want **duo** to be [see how to get](#Duo-moods). | `false` |

#### Details

- If any mood is passed, it will use the default happy mood. Since duo is always happy

<hr>

## Duo moods

You can choose wich <strong>duo selfie</strong> you prefer by passing the following id's as a `mood` parameter.

> [!NOTE]
> All Duo beautiful selfies were downloaded from Duolingo design website, if interested you can see more about [here](https://design.duolingo.com/).

<table>
  <tr>
     <td>greeting</td>
     <td>passionate</td>
     <td>sad</td>
     <td>whistler</td>
  </tr>
  <tr>
    <td><img src="/assets/images/d1.jpg" width=150></td>
    <td><img src="/assets/images/d2.jpg" width=150></td>
    <td><img src="/assets/images/d3.jpg" width=150></td>
    <td><img src="/assets/images/d4.jpg" width=150></td>
  </tr>
 </table>


 <details>
  <summary>See all</summary>
  <table>
  <tr>
     <td>angry</td>
     <td>chasing</td>
     <td>splash</td>
     <td>flirting</td>
  </tr>
  <tr>
    <td><img src="/assets/images/d5.jpg" width=150></td>
    <td><img src="/assets/images/d6.jpg" width=150></td>
    <td><img src="/assets/images/d7.jpg" width=150></td>
    <td><img src="/assets/images/d8.jpg" width=150></td>
  </tr>
 </table>
 <table>
  <tr>
     <td>happy</td>
     <td>delighted</td>
     <td>splish</td>
     <td>bored</td>
  </tr>
  <tr>
    <td><img src="/assets/images/d9.jpg" width=150></td>
    <td><img src="/assets/images/d10.jpg" width=150></td>
    <td><img src="/assets/images/d11.jpg" width=150></td>
    <td><img src="/assets/images/d12.jpg" width=150></td>
  </tr>
 </table>
  <table>
  <tr>
     <td>struggling</td>
  </tr>
  <tr>
    <td><img src="/assets/images/d13.jpg" width=150></td>
  </tr>
 </table>
</details>

<br>
 

 > [!IMPORTANT] 
> I am not a graphic designer, for more that it can look like.

<hr>

## UserID

This is the most tricky part but should'nt be hard, i was not able to find an easy way to get your own user ID so we will have to scrap it by ourselfs. Don't worry it is really simple :).

#### step 1
 - <strong>Log-in</strong> at [duolingo](https://www.duolingo.com) official website with your own account.
 - <strong>Open dev tools</strong> by pressing `F12` or click with right button and selecting `inspect`.
 - Select the <strong>network tab</strong>.

#### step 2

Here you will filter the request by `user`. 
<img src="/assets/screenshots/filter.png">

And see if you find something related to user ID. The number in `user` field is your userID :D.
<img src="/assets/screenshots/request.png">


> [!WARNING]
> I am really not sure if this is a sensytive information :/. That is my ID btw.